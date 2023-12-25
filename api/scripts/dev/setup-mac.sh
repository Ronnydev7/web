install_if_not_installed() {
    local software_name="$1"
    local installation_command="$2"

    if ! command -v "$software_name" &> /dev/null; then
        echo "$software_name not found. Installing..."
        eval "$installation_command"
    else
        echo "$software_name is already installed."
    fi
}

# Setup Atlas
install_if_not_installed "atlas" "curl -sSf https://atlasgo.sh | sh"