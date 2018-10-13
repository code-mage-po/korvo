rm -rf dist
cross-env NODE_ENV=production webpack --progress --hide-modules
rm -rf ../static_pages/note-app
mkdir -p ../static_pages/note-app
cp -f index.html ../static_pages/note-app/
cp -rf dist ../static_pages/note-app/
