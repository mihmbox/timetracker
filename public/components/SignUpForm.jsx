import React from 'react';
import jQuery from 'jquery';

var errorCodes = {
    EmailExistsOrInvalid: 1,
    PasswordIsWeak: 2,
    PasswordsDontMatchErrorCode: 3
};
export default class SignUpForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEmailInvalid: this.props.errorCode == 1,
            isPasswordWeak: this.props.errorCode == 2,
            isPasswordDoesntMatch: this.props.errorCode == 3
        };
    }

    static defaultProps = {
        email: '',
        errorCode: 0
    };

    static propTypes = {
        email: React.PropTypes.string,
        errorCode: React.PropTypes.number
    };

    componentDidMount() {
        //validateForm();
    }

    validateEmail = () => {
        clearTimeout(this.emailValidationTimeout);

        var el = React.findDOMNode(this.refs.email);
        if (!el.checkValidity()) {
            this.setState({
                isEmailInvalid: true
            });
            return;
        }

        this.emailValidationTimeout = setTimeout(() => {
            jQuery.ajax({
                url: '/api/signup/validate_email/' + el.value
            }).done(() => {
                this.setState({
                    isEmailInvalid: false
                });
            }).fail(() => {
                this.setState({
                    isEmailInvalid: true
                });
            });
        }, 700);
    };

    validatePassword = () => {
        var password = React.findDOMNode(this.refs.password).value.trim();
        this.setState({
            isPasswordWeak: password.length < 6
        });
        this.comparePasswords();
    };

    comparePasswords = () => {
        var password = React.findDOMNode(this.refs.password).value.trim();
        var passwordConfirm = React.findDOMNode(this.refs.passwordConfirm).value.trim();
        this.setState({
            isPasswordDoesntMatch: password != passwordConfirm
        });
    };

    // Focuses on specified ref element
    focus = (ref) => {
        var _this = this;
        return function(e) {
            React.findDOMNode(_this.refs[ref]).focus();
        }
    };

    render() {
        var errorIcon = <i className="SignUpForm_erroricon material-icons prefix red-text">warning</i>;

        return (
            <div className="SignUpForm" ref="test">
                <div>
                    <div className="SignUpForm_header card-panel teal white-text"><span>Sign Up</span></div>
                </div>

                <form className="SignUpForm_form col s12 white" method="post" action="/signup">
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="email" name="email" type="email"
                                   className={"novalidate " +  (this.state.isEmailInvalid ? 'invalid' : '')}
                                   defaultValue={this.props.email}
                                   onChange={this.validateEmail}
                                   ref="email"/>
                            <label for="email"
                                   onClick={this.focus('email')}
                                   data-error="Email is invalid or already taken"
                                   data-success="Right">Email</label>
                            {this.state.isEmailInvalid && errorIcon}
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password" name="password" type="password"
                                   className={"novalidate " +  (this.state.isPasswordWeak ? 'invalid': '')}
                                   ref="password"
                                   onChange={this.validatePassword}
                                />
                            <label for="password"
                                   onClick={this.focus('password')}
                                   data-error="Password is weak (minimum is 6 characters)"
                                   data-success="">Password</label>
                            {this.state.isPasswordWeak && errorIcon}
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password-confirm" name="passwordconfirm" type="password"
                                   className={"novalidate " +  (this.state.isPasswordDoesntMatch ? 'invalid': '')}
                                   ref="passwordConfirm"
                                   onChange={this.comparePasswords}/>
                            <label for="password-confirm"
                                   onClick={this.focus('passwordConfirm')}
                                   data-error="Password doesn't match"
                                   data-success="">Confirm password</label>
                            {this.state.isPasswordDoesntMatch && errorIcon}
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