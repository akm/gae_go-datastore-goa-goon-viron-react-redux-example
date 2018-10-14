import { call, takeEvery, put } from 'redux-saga/effects'
import { Action } from 'typescript-fsa';

import { MemosApiFactory, Memo, MemoCollection } from '../api/'
import { addMemoAction, refreshPostAction } from '../actions/index'

export default function* rootSaga() {
  yield takeEvery("REFRESH_ASYNC", refreshAsync)
  yield takeEvery("SAVE_ASYNC", saveAsync)
}

function* refreshAsync() {
  const memos = yield call(fetchMemos); // TODO handle error
  yield put(refreshPostAction(memos))
}

function fetchMemos() {
  return MemosApiFactory().memosList()
    .then((memos: MemoCollection) => {
      console.log("SUCCESS response", memos);
      return memos;
    })
    .catch((err) => {
      console.log("ERROR", err)
      return err
    });
}

function* saveAsync(action: Action<{content: string}>) {
  const memo = yield call(postMemo, action); // TODO handle error
  yield put(addMemoAction(memo))
}

function postMemo(action: Action<{content: string}>) {
  return MemosApiFactory().memosCreate(action.payload)
    .then((memo: Memo) => {
      console.log("SUCCESS response", memo)
      return memo;
    })
    .catch((err) => {
      console.log("ERROR ", err);
      return err;
    })
}
