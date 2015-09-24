import React from 'react';

export default class SignUpForm extends React.Component {
    render() {
        return (
            <div className="SignUpForm">
                {this.props.showHeader && (
                    <div>
                        <div className="SignUpForm_header card-panel teal white-text"><span>Sign Up</span></div>
                    </div>
                )}
                <form className="SignUpForm_form col s12 white" method="post" action="/signup">
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="email" type="email" className="validate"/>
                            <label for="email" data-error="Email is invalid or already taken"
                                   data-success="right">Email</label>
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <input id="password" type="password" className="validate"/>
                            <label for="password"
                                   data-error="Password can't be blank and is too short (minimum is 7 characters)"
                                   data-success="right">Password</label>
                        </div>
                    </div>
                    <div className="row">
                        <button type="submit" className="waves-effect  btn yellow darken-4 col l12">Sign up</button>
                    </div>
                </form>
            </div>
        );
    }
}
SignUpForm.defaultProps = {
    showHeader: true
};
SignUpForm.propTypes = {
    showHeader: React.PropTypes.bool
};