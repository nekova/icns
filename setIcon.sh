ICON=$1
DIR=$2

ICONFILE=$'Icon\r'
RSRCFILE="${ICONFILE}"

cd "${DIR}"
touch "${RSRCFILE}"
cat << EOF > "${RSRCFILE}"
read 'icns' (-16455) "${ICON}";
EOF

Rez "${RSRCFILE}" -o "${ICONFILE}"
if [ "x${RSRCFILE}" = "x${ICONFILE}" ]; then
  SetFile -a C .
fi
