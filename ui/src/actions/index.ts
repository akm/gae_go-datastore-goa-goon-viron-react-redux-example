import actionCreatorFactory from 'typescript-fsa';
import { Memo, MemoPayload } from '../api'

const actionCreator = actionCreatorFactory();

export const refreshPostAction = actionCreator<Memo[]>('REFRESH_POST');
export const addMemoAction = actionCreator<Memo>('ADD_MEMO');
export const saveAsyncAction = actionCreator<MemoPayload>('SAVE_ASYNC');
export const refreshAsyncAction = actionCreator('REFRESH_ASYNC');
