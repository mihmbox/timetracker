'use strict';

import React from 'react';
import SignInForm from './components/SignInForm.jsx';

React.render(<SignInForm showHeader={false} />, document.getElementById('signin-form'));

$(function() {
    $('.parallax').parallax();
    $('.button-collapse').sideNav();
});
