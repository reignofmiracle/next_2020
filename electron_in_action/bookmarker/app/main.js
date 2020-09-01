const { app, BrowserWindow } = require('electron');

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true
    }
  })

  win.loadFile("./app/index.html")
}

app.whenReady().then(createWindow)

// let mainWindow = null; // #A

// app.on('ready', () => {
//   console.log('Hello from Electron.');
//   mainWindow = new BrowserWindow();
//   mainWindow.webContents.loadURL(`file://${__dirname}/index.html`); // #A
// });

