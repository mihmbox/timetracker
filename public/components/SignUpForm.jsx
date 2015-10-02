import React from 'react';
import jQuery from 'jquery';

export default class SignUpForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEmailValid: true,
            isPasswordValid: true,
            isPasswordConfirmValid: true
        };
    }

    componentDidMount() {
        //validateForm();
    }

    validateEmail = () => {
        clearTimeout(this.emailValidationTimeout);

        var el = React.findDOMNode(this.refs.email);
        if (!el.checkValidity()) {
            this.setState({
                isEmailValid: false
            });
            return;
        }

        this.emailValidationTimeout = setTimeout(() => {
            jQuery.ajax({
                url: '/api/signup/validate_email/' + el.value
            }).done(() => {
                this.setState({
                    isEmailValid: true
                });
            }).fail(() => {
                this.setState({
                    isEmailValid: false
                });
            });
        }, 700);
    };

    render() {
        var errorIcon = <i className="SignUpForm_erroricon material-icons prefix red-text">warning</i>;
        //<i className="SignUpForm_okicon material-icons prefix active">done</i>
        return (
            <div className="SignUpForm" ref="test">
                <div>
                    <div className="SignUpForm_header card-panel teal white-text"><span>Sign Up</span></div>
                </div>

                <form className="SignUpForm_form col s12 white" method="post" action="/signup">
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="email" name="email" type="email"
                                   className={"validate " +  (!this.state.isEmailValid && 'invalid')}
                                   onChange={this.validateEmail} ref="email"/>
                            <label for="email" data-error="Email is invalid or already taken"
                                   data-success="Right">Email</label>
                            {!this.state.isEmailValid && errorIcon}
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password" name="password" type="password" className="validate"/>
                            <label for="password"
                                   data-error="Password can't be blank and is too short (minimum is 7 characters)"
                                   data-success="Correct">Password</label>
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password-confirm" name="passwordconfirm" type="password" className="validate"/>
                            <label for="password-confirm"
                                   data-error="Passwords don't match"
                                   data-success="Correct">Confirm password</label>
                        </div>
                    </div>
                    <div className="row">
                        <button type="submit" className="waves-effect  btn yellow darken-4 col l12">
                            Create an account
                        </button>
                    </div>
                </form>
            </div>
        );
    }
}
SignUpForm.defaultProps = {
    //showHeader: true,
    //loginFailed: false
};
SignUpForm.propTypes = {
    //showHeader: React.PropTypes.bool,
    //loginFailed: React.PropTypes.bool
};