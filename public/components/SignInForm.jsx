import React from 'react';

export default class SignInForm extends React.Component {
    //constructor(props) {
    //    super(props);
    //    this.state = {
    //        isFormValid: false
    //    };
    //}

    render() {
        return (
            <div className="SignInForm">
                {this.props.showHeader && (
                    <div>
                        <div className="SignInForm_header card-panel teal white-text"><span>Sign In</span></div>
                    </div>
                )}
                <form className="SignInForm_form col s12 white" method="post" action="/signin">
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="email" type="email" className="validate" />
                            <label for="email" data-error="Email is invalid or already taken" data-success="Right">Email</label>
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password" type="password" className="validate"/>
                            <label for="password"
                                   data-error="Password can't be blank and is too short (minimum is 7 characters)"
                                   data-success="Right">Password</label>
                        </div>
                    </div>
                    {this.props.loginFailed && (
                        <div className="row SignInForm_failed red-text">
                            <span>Incorrect email or password.</span>
                        </div>
                    )}
                    <div className="row">
                        <button type="submit" className="waves-effect  btn yellow darken-4 col l12">
                            Sign in
                        </button>
                    </div>
                </form>
            </div>
        );
    }
}
SignInForm.defaultProps = {
    showHeader: true,
    loginFailed: false
};
SignInForm.propTypes = {
    showHeader: React.PropTypes.bool,
    loginFailed: React.PropTypes.bool
};