const fs = require('fs');
const path = require('path');

const appName = process.argv[2];
if (!appName) {
  console.error('❌ Please provide an app name: npm run remove:app my-app');
  process.exit(1);
}

const appPath = path.join(__dirname, '../apps', appName);

// Check if the app exists
if (!fs.existsSync(appPath)) {
  console.error(`❌ App "${appName}" does not exist.`);
  process.exit(1);
}

// Delete the app directory
fs.rmSync(appPath, { recursive: true, force: true });
console.log(`✅ App "${appName}" removed successfully.`);
