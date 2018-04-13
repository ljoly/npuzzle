import React, {Component} from 'react';
import TileSection from './tiles/TileSection.jsx'
import ButtonSection from './buttons/ButtonSection.jsx'
import io from 'socket.io-client'

class App extends Component{
    constructor(props){
        super(props);
        var socket = io("http://localhost:3000")
        socket.on('connect', function(){
            console.log(socket)
            console.log("CONNECTED TO SERVER")            
        })
        this.state = {
            puzzle: [1, 2, 3, 4, 5, 16, 17, 18, 19, 6, 15, 24, 0, 20, 7 ,14, 23, 22, 21, 8, 13, 12, 11, 10, 9],
            size: 5,
            socket
        }
    }
    socketHandler(){
        this.state.socket.on('state', function(data){
            console.log('State', data)
            // this.setState({puzzle: data});
        }.bind(this))
    }
    prevState(){
        console.log("okok")
        this.state.socket.emit('prevState')
    }
    nextState(){
        this.state.socket.emit('nextState')
    }
    go(){
        this.state.socket.emit('go')
    }
    render(){
        this.socketHandler()
        return (
            <div>
                <TileSection
                    {...this.state}
                />
                <ButtonSection
                    prevState = {this.prevState.bind(this)}
                    nextState = {this.nextState.bind(this)}
                    go = {this.go.bind(this)}
                    {...this.state}
                />
            </div>
        )
    }
}

export default App