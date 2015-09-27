'use strict';

import React from 'react'
import SignInForm from '../../components/SignInForm.jsx'

export default class SignIn extends React.Component {
    render() {
        return (
            <div className="row valign-wrapper">
                <div className="valign col l3 offset-l4 m6 offset-m3 s12">
                    <SignInForm showHeader={true} />
                </div>
            </div>
        );
    }
}