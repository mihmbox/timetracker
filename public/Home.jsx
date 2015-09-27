'use strict';

import React from 'react'
import {Router, Route, Link, IndexRoute, Redirect} from 'react-router'
import { createHistory, createHashHistory } from 'history';

import Home from './views/home/index.jsx'
import SignIn from './views/home/signin.jsx'

class HomeApp extends React.Component {
    render() {
        return (
            <div>
                {this.props.children}
            </div>
        );
    }
}

//const history = createHashHistory();
const history = createHistory();
React.render((
    <Router history={history}>
        <Route path="/" component={HomeApp}>
            <IndexRoute component={Home}/>
            <Route path="signin" component={SignIn}/>
        </Route>
    </Router>
), document.getElementById('HomeApp'));

