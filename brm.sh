# !/bin/bash

# 定义 Homebrew 源地址
OFFICIAL_BREW_REPO="https://github.com/Homebrew/brew.git"
OFFICIAL_CORE_REPO="https://github.com/Homebrew/homebrew-core.git"

ALIYUN_BREW_REPO="https://mirrors.aliyun.com/homebrew/brew.git"
ALIYUN_CORE_REPO="https://mirrors.aliyun.com/homebrew/homebrew-core.git"
ALIYUN_BOTTLES_DOMAIN="https://mirrors.aliyun.com/homebrew/homebrew-bottles"

TSINGHUA_BREW_REPO="https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"
TSINGHUA_CORE_REPO="https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git"
TSINGHUA_BOTTLES_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"

USTC_BREW_REPO="https://mirrors.ustc.edu.cn/brew.git"
USTC_CORE_REPO="https://mirrors.ustc.edu.cn/homebrew-core.git"
USTC_BOTTLES_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"

# 定义设置源的函数
function set_homebrew_source {
    echo "正在设置 Homebrew Git 仓库源为: $1"
    git -C "$(brew --repo)" remote set-url origin "$1"

    echo "正在设置 Homebrew Core 仓库源为: $2"
    git -C "$(brew --repo homebrew/core)" remote set-url origin "$2"

    # 设置或更新 HOMEBREW_BOTTLE_DOMAIN（如果需要）
    if [[ -n "$3" ]]; then
        update_shell_config "$3"
    else
        remove_bottle_domain_from_shell_config
    fi

    display_current_sources
}

# 更新 shell 配置文件
function update_shell_config {
    local config_file
    if [[ -f "$HOME/.zshrc" ]]; then
        config_file="$HOME/.zshrc"
    elif [[ -f "$HOME/.bashrc" ]]; then
        config_file="$HOME/.bashrc"
    else
        echo "未找到 shell 配置文件。"
        return
    fi

    if grep -q "export HOMEBREW_BOTTLE_DOMAIN" "$config_file"; then
        sed -i '' "s|export HOMEBREW_BOTTLE_DOMAIN=.*|export HOMEBREW_BOTTLE_DOMAIN=$1|" "$config_file"
    else
        echo "export HOMEBREW_BOTTLE_DOMAIN=$1" >> "$config_file"
    fi
}

function remove_bottle_domain_from_shell_config {
    sed -i '' '/export HOMEBREW_BOTTLE_DOMAIN/d' "${HOME}/.zshrc" "${HOME}/.bashrc"
}

function display_current_sources {
    echo "\n当前使用的 Homebrew Git 仓库源:"
    git -C "$(brew --repo)" remote -v

    echo "\n当前使用的 Homebrew Core 仓库源:"
    git -C "$(brew --repo homebrew/core)" remote -v
}

# 定义测试源速度的函数
function test_source_speed {
    local source_name=$2
    # echo "正在测试源：${source_name}"
    
    local start_time=$(gdate +%s%N)
    
    # 使用 git ls-remote 测试速度
    git ls-remote -h "$1" HEAD > /dev/null 2>&1

    local end_time=$(gdate +%s%N)
    local elapsed=$(echo "scale=3; ($end_time - $start_time)/1000000" | bc)
    echo "正在测试源：${source_name} 耗时: ${elapsed} 毫秒"
}

function test_all_sources_speed {
    echo "\n测试各个源的速度:"
    test_source_speed "$OFFICIAL_BREW_REPO" "官方源"
    test_source_speed "$ALIYUN_BREW_REPO" "阿里云"
    test_source_speed "$TSINGHUA_BREW_REPO" "清华大学"
    test_source_speed "$USTC_BREW_REPO" "中科大"
}

# 主菜单
echo "请选择 Homebrew 源："
echo "1. 官方源"
echo "2. 阿里云"
echo "3. 清华大学"
echo "4. 中科大"
echo "5. 测试所有源的速度"
read -p "请输入选择 [1-5]: " choice

case $choice in
    1)
        set_homebrew_source "$OFFICIAL_BREW_REPO" "$OFFICIAL_CORE_REPO"
        ;;
    2)
        set_homebrew_source "$ALIYUN_BREW_REPO" "$ALIYUN_CORE_REPO" "$ALIYUN_BOTTLES_DOMAIN"
        ;;
    3)
        set_homebrew_source "$TSINGHUA_BREW_REPO" "$TSINGHUA_CORE_REPO" "$TSINGHUA_BOTTLES_DOMAIN"
        ;;
    4)
        set_homebrew_source "$USTC_BREW_REPO" "$USTC_CORE_REPO" "$USTC_BOTTLES_DOMAIN"
        ;;
    5)
        test_all_sources_speed
        ;;
    *)
        echo "无效的选择。"
        exit 1
        ;;
esac

echo "\nHomebrew 源切换完成。"
