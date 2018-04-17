import React, { Component } from 'react';
import PropTypes from 'prop-types'

class ButtonSection extends Component {
    onPrev() {
        this.props.prevState();
    }
    onGo() {
        this.props.go();
    }
    onNext() {
        this.props.nextState();
    }
    render() {
        return (
            <div>
                <div style={{ textAlign: "center" }}>
                    <button onClick={this.onPrev.bind(this)}>
                        PREV
                    </button>
                    <button onClick={this.onGo.bind(this)}>
                        GO
                    </button>
                    <button onClick={this.onNext.bind(this)}>
                        NEXT
                    </button>
                </div>
            </div>
        )
    }
}

ButtonSection.propTypes = {
    prevState: PropTypes.func.isRequired,
    nextState: PropTypes.func.isRequired,
    go: PropTypes.func.isRequired
}

export default ButtonSection