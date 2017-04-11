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
        console.log(HelperApp);
        debugger;
        HelperApp.initSW()
        .then(res => {
            console.log(res);
        })
        .catch(e => {
            console.log(e);
        });
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
            </div>
        )
        
    }
}