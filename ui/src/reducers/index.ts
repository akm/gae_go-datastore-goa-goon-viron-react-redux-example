import { reducerWithInitialState } from "typescript-fsa-reducers";

import { addMemoAction, refreshPostAction } from '../actions/index'

export default reducerWithInitialState([])
.case(refreshPostAction, (state, memos) => memos )
.case(addMemoAction, (state, memo) => [...state, memo])
