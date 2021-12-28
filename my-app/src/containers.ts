// Reduxの値をpropsで受け取るようにする
// Redux Containerの実装
import { connect, ConnectedProps } from "react-redux";
import { clickSquare, jumpToPast } from "./action";
import { Game } from "./components";

export type StateByProps = {}
export type DispatchByProps = {
    handleClick: (index: number) => void
    jumpTo: (step: number) => void
  }

const mapState = (state: { game: any; }) : StateByProps => {
    return state.game;
}

const mapDispatch = (dispatch:any) : DispatchByProps => ({
        handleClick: (index: number) => {
            dispatch(clickSquare(index));
        },
        jumpTo: (step: number) => {
            dispatch(jumpToPast(step));
        }
});

const connector = connect(mapState, mapDispatch)
export type PropsFromRedux = ConnectedProps<typeof connector>;
export default connector(Game);