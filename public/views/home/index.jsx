'use strict';

import React from 'react'
import Link from 'react-router'
import SignInForm from '../../components/SignInForm.jsx'

export default class Home extends React.Component {
    componentDidMount() {
        $(function () {
            $('.parallax').parallax();
            $('.button-collapse').sideNav();
        });
    }

    render() {
        return (
            <div>
                <div id="index-banner" className="parallax-container valign-wrapper">
                    <div className="container row no-pad-bot valign">
                        <div className="col m9 s12">
                            <h1 className="header center teal-text white-text">Track your time</h1>

                            <div className="row center">
                                <Link to="/signin" className="btn-large waves-effect waves-light teal lighten-1">Get Started</Link>
                            </div>
                        </div>
                        <div className="col m3">
                            <SignInForm showHeader={false} />
                        </div>
                    </div>
                    <div className="parallax">
                        <img src="/public/img/home_banner.png" alt="Track your time"/>
                    </div>
                </div>
                <div className="container">
                    <div className="section">
                        {/*   Icon Section   */}
                        <div className="row">
                            <div className="col s12 m4">
                                <div className="icon-block">
                                    <h2 className="center brown-text"><i className="material-icons">flash_on</i></h2>
                                    <h5 className="center">Speeds up development</h5>

                                    <p className="light">
                                        Nuptia cadunts, tanquam grandis navis.
                                        Nuptia cadunts, tanquam grandis navis.
                                        Nuptia cadunts, tanquam grandis navis.
                                    </p>
                                </div>
                            </div>
                            <div className="col s12 m4">
                                <div className="icon-block">
                                    <h2 className="center brown-text"><i className="material-icons">group</i></h2>
                                    <h5 className="center">User Experience Focused</h5>

                                    <p className="light">
                                        A falsis, clabulare varius historia.
                                        A falsis, clabulare varius historia.
                                        A falsis, clabulare varius historia.
                                        A falsis, clabulare varius historia.
                                    </p>
                                </div>
                            </div>
                            <div className="col s12 m4">
                                <div className="icon-block">
                                    <h2 className="center brown-text"><i className="material-icons">settings</i></h2>
                                    <h5 className="center">Easy to work with</h5>

                                    <p className="light">
                                        Done by developers for develoers.
                                        The scurvy scabbard roughly breaks the bilge rat.
                                        The scurvy scabbard roughly breaks the bilge rat.
                                        The scurvy scabbard roughly breaks the bilge rat.
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="divider"/>
                    <div className="section">
                        <div className="row">
                            <div className="col l9 s12">
                                <h5 className>Company Bio</h5>

                                <p className>
                                    We are a team of ......
                                </p>
                            </div>
                            <div className="col l3 s12">
                                <h5 className>Connect</h5>
                                <ul>
                                    <li><a className href="#!">Link 1</a></li>
                                    <li><a className href="#!">Link 2</a></li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}



//
//import SignInForm from './components/SignInForm.jsx';
//React.render(<SignInForm showHeader={false} />, document.getElementById('signin-form'));
//
//$(function () {
//    $('.parallax').parallax();
//    $('.button-collapse').sideNav();
//});