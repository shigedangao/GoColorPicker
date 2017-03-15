const { app, BrowserWindow } = require('electron');
const path = require('path');
const url = require('url');

// Global variable 
let win; 

(() => {
    
    /**
     * Create Window 
     *      Create the window 
     */
    const createWindow = () => {
        win = new BrowserWindow({width: 800, height: 600});
        win.loadURL('http://localhost:1698/');
        attachedEvent('closed', win, () => {wind = null}, null);
        win.webContents.openDevTools();
    };

    /**
     * Attached Event 
     *      Add event to the window when need
     * @param {String} event
     * @param {Function} callback
     * @param {Mixed} params
     */
    const attachedEvent = (event, bindVar, callback, params = null) => {
        if (params !== undefined || params !== null)
            callback.bind(null, params);

        // append an event and add it's callback 
        bindVar.on(event, callback)
    };


    attachedEvent('ready', app, createWindow, null);
    attachedEvent('window-all-closed', app, () => {
        if (process.platform !== 'darwin')
            app.quit();
    }, null);
    attachedEvent('activate', app, () => {
        if (win === null)
            createWindow();
    }, null);
})();