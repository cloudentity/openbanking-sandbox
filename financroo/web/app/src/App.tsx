import React, {useEffect, useState} from 'react';
import './App.css';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import {Switch} from 'react-router';
import {createMuiTheme, ThemeProvider} from "@material-ui/core";
import {StylesProvider} from '@material-ui/core/styles';
import superagent from 'superagent';
import Progress from "./components/Progress";
import {toJson} from "./api/api-base";
import Register from "./components/Register";
import Main from "./components/Main";
import {ReactQueryDevtools} from "react-query-devtools";

const theme = createMuiTheme({
  palette: {
    primary: {
      main: "#36C6AF"
    }
  },
  overrides: {
    MuiTableRow: {
      root: {
        '&$selected': {
          backgroundColor: "rgba(54, 198, 175, 0.08)",
          '&:hover': {
            backgroundColor: "rgba(54, 198, 175, 0.2)",
          }
        }
      }
    },
    MuiTableCell: {
      root: {
        borderBottom: 'none'
      }
    }
  }
});

export type Config = {}

function App() {
  const [progress, setProgress] = useState(true);
  const [config, setConfig] = useState<Config | null>(null);

  useEffect(() => {
    superagent
      .get('/config.json')
      .then(toJson)
      .then(res => setConfig(res))
      .finally(() => setProgress(false));
  }, []);

  return (
    <>
      <ReactQueryDevtools/>
      <ThemeProvider theme={theme}>
        <StylesProvider injectFirst>
          {progress && <Progress/>}
          {!progress && (
            <Router>
              <Switch>
                <Route exact path={"/register"} render={() => <Register/>}/>
                <Route exact path={"/"} render={() => <Main/>}/>
              </Switch>
            </Router>
          )}
        </StylesProvider>
      </ThemeProvider>
    </>
  );
}

export default App;
