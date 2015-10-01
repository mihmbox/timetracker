'use strict';

import React from 'react';
import SignUpForm from './components/SignUpForm.jsx';

React.render(
    (
        <div className="row valign-wrapper">
            <div className="valign col l3 offset-l4 m6 offset-m3 s12">
                <SignUpForm />
            </div>
        </div>
    ), document.getElementById('SignupApp')
);