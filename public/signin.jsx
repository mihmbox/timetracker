'use strict';

import React from 'react';
import SignInForm from './components/SignInForm.jsx';

React.render(
    (
        <div className="row valign-wrapper">
            <div className="valign col l3 offset-l4 m6 offset-m3 s12">
                <SignInForm loginFailed={SigninAppData.failed} email={SigninAppData.email}/>
            </div>
        </div>
    ), document.getElementById('SigninApp')
);