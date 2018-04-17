import React, { Component } from 'react';
import TileSection from './tiles/TileSection.jsx'
import ButtonSection from './buttons/ButtonSection.jsx'
import io from 'socket.io-client'
import cron from 'node-cron'

class App extends Component {
    constructor(props) {
        super(props);
        const socket = io("http://localhost:3000")
        const task = cron.schedule('* * * * * *', function () {
            this.goNextState()
        }.bind(this))
        task.stop()
        this.state = {
            index: 0,
            all: [],
            puzzle: [],
            size: 0,
            socket,
            task
        }
    }
    socketHandler() {
        this.state.socket.on('connect', function () {
        })
        this.state.socket.on('state', function (data) {
            data = JSON.parse(data)
            this.setState({ all: data });
            this.setState({ puzzle: this.state.all[0].Board, size: this.state.all[0].Size })
        }.bind(this))
    }
    prevState() {
        this.state.task.stop()
        if (this.state.index > 0) {
            this.state.index--
            this.setState({ puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size })
        }
    }
    nextState() {
        this.state.task.stop()
        if (this.state.index < this.state.all.length - 1) {
            this.state.index++
            this.setState({ puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size })
        }
    }
    goNextState() {
        if (this.state.index < this.state.all.length - 1) {
            this.state.index++
            this.setState({ puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size })
        }
    }
    go() {
        this.state.task.start()
        if (this.state.index === this.state.all.length - 1) {
            this.state.task.stop()
        }
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