create_directory() {
  local folder="$1"

  # Check if the folder exists, and create it if not
  if [ ! -d "$folder" ]; then
    mkdir -p "$folder"
  fi
}

echo "++++++++++++++++++++"
echo "installing i18n to app"

echo "Installing intl..."
rm -rf ../app/src/generated/intl/
create_directory ../app/src/generated
cp -r intl ../app/src/generated/intl/

echo "Installing translations..."
rm -rf ../app/src/generated/translations/
create_directory ../app/src/generated
cp -r translations ../app/src/generated/translations/
echo "Successfully installed i18n to app"
echo "---------------------"

echo "++++++++++++++++++++"
echo "installing i18n to api"
cd ../api
go generate ./generators/i18n
cd -
echo "Successfully installed i18n to api"
echo "---------------------"