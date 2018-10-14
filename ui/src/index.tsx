import * as React from 'react';
import * as ReactDom from 'react-dom';
import Root from './components/Root';

import configureStore from './store';
import { saveAsyncAction, refreshAsyncAction } from './actions/index'
import './styles/common.scss'

const store = configureStore([])

const render = () => ReactDom.render(
    <Root
      memos={store.getState()}
      onRefresh={ () => store.dispatch(refreshAsyncAction()) }
      onSave={ (content) => store.dispatch(saveAsyncAction({content: content})) }
    />,
    document.getElementById('root'),
);

render()
store.subscribe(render)
