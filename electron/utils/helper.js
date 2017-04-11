// Create an empty Object with the react props inside 
const _pickerHelper = Object.create({});

/**
 * Init SW 
 *          Init the service worker. The service worker will only be use for handling the 
 *          transiting datas and the connection between the app and the Go socket.
 * @public
 */
_pickerHelper.initSW = () => {
    return new Promise((resolve, reject) => {
         if ('serviceWorker' in navigator){
            navigator.serviceWorker.register('./sw/sw.js')
            .catch(e => {
                reject(e);
            });

            resolve(true);
         }
        else {
            resolve('sw is not supported by your webview');
        }
    });
};

export {_pickerHelper};


