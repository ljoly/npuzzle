import React, { Component } from 'react';
import TileSection from './tiles/TileSection.jsx'
import ButtonSection from './buttons/ButtonSection.jsx'
import io from 'socket.io-client'

class App extends Component {
    constructor(props) {
        super(props);
        var socket = io("http://localhost:3000")
        socket.on('connect', function () {
            console.log("CONNECTED TO SERVER")
        })
        this.state = {
            index: 0,
            all: [],
            puzzle: [],
            size: 0,
            socket
        }
    }
    socketHandler() {
        this.state.socket.on('state', function (data) {
            data = JSON.parse(data)
            this.setState({ all: data });
            this.setState({ puzzle: this.state.all[0].Board, size: this.state.all[0].Size })
        }.bind(this))
    }
    prevState() {
        if (this.state.index > 0) {
            this.state.index--
            this.setState({ puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size })
            console.log('index', this.state.index)
        }
    }
    nextState() {
        if (this.state.index < this.state.all.length - 1) {
            this.state.index++
            this.setState({ puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size })
            console.log('index', this.state.index)
        }
    }
    go() {
        var i = this.state.index
        const self = this
        while (i < self.state.all.length - 1) {
            setInterval(function () {
                setTimeout(() => self.setState({ puzzle: self.state.all[i].Board }), 1000)
            }.bind(self), 1000)
            i++
        }
        this.state.index = i
        console.log('index', this.state.index)
    }
    componentWillMount() {
        this.socketHandler()
    }
    render() {
        return (
            <div>
                <TileSection
                    {...this.state}
                />
                <ButtonSection
                    prevState={this.prevState.bind(this)}
                    nextState={this.nextState.bind(this)}
                    go={this.go.bind(this)}
                    {...this.state}
                />
            </div>
        )
    }
}

export default App