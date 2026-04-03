#!/bin/bash
set -eo pipefail

# 将本项目 skills/ 目录下的 SKILL.md 安装到本机所有已检测到的 AI Agent 环境
# 用法:
#   ./scripts/install-skills.sh              安装到所有检测到的 Agent
#   ./scripts/install-skills.sh --uninstall   从所有 Agent 中卸载
#   ./scripts/install-skills.sh --list        列出支持的 Agent 及检测结果

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
SKILLS_SRC="$PROJECT_DIR/skills"

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
DIM='\033[2m'
NC='\033[0m'

# ──────────────────────────────────────────────────────────
# Agent 路径映射表（对齐 vercel-labs/skills agents.ts）
# 格式: "agent_name|path"
# ──────────────────────────────────────────────────────────
AGENTS=(
    "adal|$HOME/.adal/skills"
    "amp|$HOME/.config/agents/skills"
    "antigravity|$HOME/.gemini/antigravity/skills"
    "augment|$HOME/.augment/skills"
    "bob|$HOME/.bob/skills"
    "claude-code|$HOME/.claude/skills"
    "cline|$HOME/.agents/skills"
    "codebuddy|$HOME/.codebuddy/skills"
    "codex|$HOME/.codex/skills"
    "command-code|$HOME/.commandcode/skills"
    "continue|$HOME/.continue/skills"
    "cortex|$HOME/.snowflake/cortex/skills"
    "crush|$HOME/.config/crush/skills"
    "cursor|$HOME/.cursor/skills"
    "deepagents|$HOME/.deepagents/agent/skills"
    "droid|$HOME/.factory/skills"
    "firebender|$HOME/.firebender/skills"
    "gemini-cli|$HOME/.gemini/skills"
    "github-copilot|$HOME/.copilot/skills"
    "goose|$HOME/.config/goose/skills"
    "iflow-cli|$HOME/.iflow/skills"
    "junie|$HOME/.junie/skills"
    "kilo|$HOME/.kilocode/skills"
    "kimi-cli|$HOME/.config/agents/skills"
    "kiro-cli|$HOME/.kiro/skills"
    "kode|$HOME/.kode/skills"
    "mcpjam|$HOME/.mcpjam/skills"
    "mistral-vibe|$HOME/.vibe/skills"
    "mux|$HOME/.mux/skills"
    "neovate|$HOME/.neovate/skills"
    "openclaw|$HOME/.openclaw/skills"
    "opencode|$HOME/.config/opencode/skills"
    "openhands|$HOME/.openhands/skills"
    "pi|$HOME/.pi/agent/skills"
    "pochi|$HOME/.pochi/skills"
    "qoder|$HOME/.qoder/skills"
    "qwen-code|$HOME/.qwen/skills"
    "roo|$HOME/.roo/skills"
    "trae|$HOME/.trae/skills"
    "trae-cn|$HOME/.trae-cn/skills"
    "warp|$HOME/.agents/skills"
    "windsurf|$HOME/.codeium/windsurf/skills"
    "zencoder|$HOME/.zencoder/skills"
)

get_name() { echo "${1%%|*}"; }
get_path() { echo "${1##*|}"; }

# ──────────────────────────────────────────────────────────
# 收集有 SKILL.md 的目录
# ──────────────────────────────────────────────────────────
collect_skills() {
    local result=""
    for dir in "$SKILLS_SRC"/*/; do
        [ -f "$dir/SKILL.md" ] && result="$result $(basename "$dir")"
    done
    echo $result
}

# ──────────────────────────────────────────────────────────
# 检测本机已安装的 Agent（父目录存在即视为已安装）
# 去重：多个 Agent 共享同一路径时只保留第一个
# ──────────────────────────────────────────────────────────
detect_agents() {
    local detected=""
    local seen_paths=""
    for entry in "${AGENTS[@]}"; do
        local name=$(get_name "$entry")
        local skills_dir=$(get_path "$entry")
        local parent_dir="$(dirname "$skills_dir")"

        # 跳过重复路径
        case " $seen_paths " in
            *" $skills_dir "*) continue ;;
        esac

        if [ -d "$parent_dir" ]; then
            detected="$detected $name"
            seen_paths="$seen_paths $skills_dir"
        fi
    done
    echo $detected
}

# 根据 agent name 查路径
agent_path() {
    local target_name="$1"
    for entry in "${AGENTS[@]}"; do
        if [ "$(get_name "$entry")" = "$target_name" ]; then
            get_path "$entry"
            return
        fi
    done
}

# ──────────────────────────────────────────────────────────
# --list：列出所有支持的 Agent 及本机检测结果
# ──────────────────────────────────────────────────────────
if [ "${1:-}" = "--list" ]; then
    echo -e "${CYAN}Supported AI Agents (${#AGENTS[@]})${NC}\n"
    printf "  %-20s %-45s %s\n" "AGENT" "PATH" "STATUS"
    printf "  %-20s %-45s %s\n" "─────" "────" "──────"
    for entry in "${AGENTS[@]}"; do
        name=$(get_name "$entry")
        skills_dir=$(get_path "$entry")
        parent_dir="$(dirname "$skills_dir")"
        display_path="${skills_dir/#$HOME/~}"
        if [ -d "$parent_dir" ]; then
            printf "  %-20s %-45s %b\n" "$name" "$display_path" "${GREEN}detected${NC}"
        else
            printf "  %-20s %-45s %b\n" "$name" "$display_path" "${DIM}not found${NC}"
        fi
    done
    exit 0
fi

# ──────────────────────────────────────────────────────────
# 检测 Agent + 收集 Skill
# ──────────────────────────────────────────────────────────
detected_agents=$(detect_agents)
skill_names=$(collect_skills)

if [ -z "$detected_agents" ]; then
    echo -e "${YELLOW}No AI agents detected on this machine.${NC}"
    echo "Run with --list to see all supported agents."
    exit 1
fi

if [ -z "$skill_names" ]; then
    echo -e "${YELLOW}No skills found.${NC} Add a SKILL.md under skills/<name>/ first."
    exit 1
fi

agent_count=$(echo $detected_agents | wc -w | tr -d ' ')
skill_count=$(echo $skill_names | wc -w | tr -d ' ')

# ──────────────────────────────────────────────────────────
# --uninstall
# ──────────────────────────────────────────────────────────
if [ "${1:-}" = "--uninstall" ]; then
    echo "Uninstalling skills from $agent_count agent(s)..."
    total=0
    for agent in $detected_agents; do
        target_base=$(agent_path "$agent")
        removed=0
        for skill_name in $skill_names; do
            target="$target_base/$skill_name"
            if [ -d "$target" ]; then
                rm -rf "$target"
                removed=$((removed + 1))
            fi
        done
        if [ $removed -gt 0 ]; then
            echo -e "  ${RED}removed${NC} $removed skill(s) from ${CYAN}$agent${NC}"
            total=$((total + removed))
        fi
    done
    echo -e "\n${GREEN}Done.${NC} Removed $total skill(s) total."
    exit 0
fi

# ──────────────────────────────────────────────────────────
# 安装
# ──────────────────────────────────────────────────────────
echo -e "Detected ${CYAN}${agent_count}${NC} agent(s): $detected_agents"
echo -e "Installing ${CYAN}${skill_count}${NC} skill(s)...\n"

total=0
for agent in $detected_agents; do
    target_base=$(agent_path "$agent")
    mkdir -p "$target_base"
    for skill_name in $skill_names; do
        src="$SKILLS_SRC/$skill_name"
        target="$target_base/$skill_name"
        rm -rf "$target"
        cp -r "$src" "$target"
        total=$((total + 1))
    done
    display_path="${target_base/#$HOME/~}"
    echo -e "  ${GREEN}installed${NC} → ${CYAN}$agent${NC} ${DIM}($display_path)${NC}"
done

echo -e "\n${GREEN}Done.${NC} Installed $skill_count skill(s) × $agent_count agent(s) = $total total."
