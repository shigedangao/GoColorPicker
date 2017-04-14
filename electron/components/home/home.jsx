import {_pickerHelper as HelperApp} from '../../utils/helper';
import React from 'react';
import ReactDOM from 'react-dom';

export default class Home extends React.Component {

    /**
     * Start
     *          Start the Service Worker
     * @public
     */
    start() {
        HelperApp.initSW()
        .then(res => {
            console.log(res);
        })
        .catch(e => {
            console.log(e);
        });
    }

    /**
     * Send Message
     */
    sendMess(){
        HelperApp.sendMessage();
    }
    /**
     * Render
     *          Render the component
     * @public
     */
    render() {
        return (
            <div>
                <button onClick={this.start}>sw</button>
                <button onClick={this.sendMess}>Message</button>
            </div>
        )
        
    }
}