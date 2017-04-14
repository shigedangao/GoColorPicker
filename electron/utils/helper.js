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
            navigator.serviceWorker.register('sw.js', {scope: './'})
            .then(reg_state => _pickerHelper.getController.call(this, reg_state))
            .then(_pickerHelper.handleReady)
            .then(_pickerHelper.listenSwMessage)
            .then(resolve(ServiceWorkerRegistration))
            .catch(e => {
                reject(e);
            });
         }
        else {
            resolve('sw is not supported by your webview');
        }
    });
};

/**
 * Listen for other events
 */
_pickerHelper.getController = (reg_state) => {
    reg_state.installing.onstatechange = () => {
        console.log(navigator.serviceWorker);
        console.log(navigator.serviceWorker.controller);
    }
}

/**
 * Listen when the service worker is ready
 */
_pickerHelper.handleReady = () => {
    navigator.serviceWorker.ready.then(reg => {
        console.log(reg);
    })
    .catch(e => {

    }); 
};

/**
 * Listen Sw Message
 */
_pickerHelper.listenSwMessage = () => {
    navigator.serviceWorker.onmessage = (e) => {
        console.log(e);
    };
};  

/**
 * Send Message
 */
_pickerHelper.sendMessage = () => {
    navigator.serviceWorker.controller.postMessage('hello');
}

export {_pickerHelper};


