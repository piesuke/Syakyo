import { render } from "react-dom";
import { createStore } from "redux";
import { app } from './reducers';
import "./index.css";
import { Provider } from 'react-redux';
import { Game } from "./components";

const store = createStore(app);

render(
  <Provider store={ store }>
    <Game />
  </Provider>,
  document.getElementById("root")
);