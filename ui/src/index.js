import React from 'react';
import ReactDom from 'react-dom';
import Root from './components/Root';

import configureStore from './store.js';

const store = configureStore([])
const action = type => store.dispatch({type})

const render = () => ReactDom.render(
    <Root
      memos={store.getState()}
      onRefresh={ () => action('REFRESH_ASYNC') }
      onSave={ (content) => store.dispatch({type: 'SAVE_ASYNC', content}) }
    />,
    document.getElementById('root'),
);

render()
store.subscribe(render)
