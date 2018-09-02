import { delay } from 'redux-saga'
import { call, takeEvery, put } from 'redux-saga/effects'

import client from '../api/client'

export default function* rootSaga() {
  yield takeEvery('REFRESH_ASYNC', refreshAsync)
  yield takeEvery('SAVE_ASYNC', saveAsync)
}

export function* refreshAsync() {
  // const action = yield take({ type: 'REFRESH_REQUEST' })
  const memos = yield call(fetchMemos);
  // const memos = [
  //   "Foo",
  //   "Bar",
  //   "Baz",
  // ];
  yield put({type: 'REFRESH_POST', memos})
}

function fetchMemos() {
  return client().listMemos("/memos")
    .then((resp) => {
      console.log("SUCCESS response", resp)
      return resp.data
      // return [
      //   "Foo",
      //   "Bar",
      //   "Baz",
      // ];
    })
    .catch( (err) => {
      console.log("ERROR err", err)
      return ['Error ' + err]
    });
}

export function* saveAsync(action) {
  const memo = yield call(postMemo, action);
  yield put({type: 'ADD_MEMO', memo})
}

function postMemo(action) {
  return client().createMemos("/memos", {content: action.content})
    .then((resp) => {
      console.log("SUCCESS response", resp)
      return resp.data
    })
    .catch( (err) => {
      console.log("ERROR err", err)
      return ['Error ' + err]
    });
}
