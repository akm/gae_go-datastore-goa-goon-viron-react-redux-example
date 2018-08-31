import React from 'react';
import ReactDom from 'react-dom';
import Root from './components/Root';

import configureStore from './store.js';

const store = configureStore([])

const render = () => ReactDom.render(
    <Root />,
    document.getElementById('root'),
);

render()
store.subscribe(render)
