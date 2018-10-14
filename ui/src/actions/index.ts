import actionCreatorFactory from 'typescript-fsa';
import { Memo } from "../states/Memo"

const actionCreator = actionCreatorFactory();

export const refreshPostAction = actionCreator<Memo[]>('REFRESH_POST');
export const addMemoAction = actionCreator<Memo>('ADD_MEMO');
export const saveAsyncAction = actionCreator<{content: string}>('SAVE_ASYNC');
export const refreshAsyncAction = actionCreator('REFRESH_ASYNC');
