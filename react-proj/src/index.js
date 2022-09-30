import * as React from 'react';
import * as ReactDOM from 'react-dom';
import App from './App';
import "regenerator-runtime/runtime";
import * as serviceWorker from './serviceWorker';
import './index.css'

window.alert = () => { };
window.prompt = () => { };
window.confirm = () => { };

ReactDOM.render(<App />, document.getElementById('root'));

serviceWorker.unregister();