import { delay } from 'redux-saga'
import { call, takeEvery, put } from 'redux-saga/effects'

import client from '../api/client'

function fetchMemos() {
  return client().listMemos("/memos")
    .then((resp) => {
      console.log("SUCCESS response", resp)
      // return resp.data
      return [
        "Foo",
        "Bar",
        "Baz",
      ];
    })
    .catch( (err) => {
      console.log("ERROR err", err)
      return ['Error ' + err]
    });
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

export default function* rootSaga() {
  yield takeEvery('REFRESH_ASYNC', refreshAsync)
}
