import {app, BrowserWindow} from 'electron';
import path from 'path';
import url from 'url'; 

//const {app, BrowserWindow} = require('electron');


(() => {
    let _window;

    /**
     * Init 
     *          Init the electron app
     * @private
     * @void
     */
    const init = function(){
        let _window = new BrowserWindow({width: 800, height: 600});
        console.log(__dirname);
        try {
            debugger;
            _window.loadURL(url.format({
                pathname: path.join(__dirname, '../index.html'),
                protocol: 'file',
                slashes: true
            }));

            // we need to recreate the window sometimes if we are using mac OS
            if (_window === null)
                init();

            _window.webContents.openDevTools();

        } catch (e) {
            console.log(e);
        }
    };


    /**
     * Add Listener 
     *          Add the listener to the app
     * 
     * @private
     * @void
     */
    const addListener = function(bindVar, event, callback){
        if (!callback || typeof event !== 'string')
            throw new Error('callback not specified');
        
        bindVar.addListener(event, callback);
    };

    /**
     * Quit 
     *          Quit the application
     */
    const quit = function(){
        if (process.platform === 'darwin')
            app.quit();
    }

    /**
     * Pre Init 
     *          Pre init the electron app
     */
    const preInit = () => {
        addListener(app, 'ready', init);
        addListener(app, 'window-all-closed', quit);
    }

    preInit();

})();


