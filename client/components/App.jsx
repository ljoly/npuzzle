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
            index: 0,
            all: [],
            // puzzle: [1, 2, 3, 4, 5, 16, 17, 18, 19, 6, 15, 24, 0, 20, 7 ,14, 23, 22, 21, 8, 13, 12, 11, 10, 9],
            puzzle: [],
            size: 0,
            socket
        }
    }
    socketHandler(){
        this.state.socket.on('state', function(data){
            console.log('State', data)
            data = JSON.parse(data)
            this.setState({all: data});
            this.setState({puzzle: this.state.all[0].Board, size: this.state.all[0].Size})
        }.bind(this))
    }
    prevState(){
        console.log('PREV')
        if (this.state.index > 0) {
            this.setState({index: this.state.index-1})
            this.setState({puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size})
        }
    }
    nextState(){
        console.log('NEXT')
        if (this.state.index < this.state.all.length) {
            this.setState({index: this.state.index+1})
            this.setState({puzzle: this.state.all[this.state.index].Board, size: this.state.all[this.state.index].Size})
        }
    }
    go(){
        var index = 0
            while (index < this.state.all.length - 1) {
                setTimeout(function(){
                index++
                console.log(index)
                console.log(this.state.all[index].Board)
                this.setState({puzzle: this.state.all[index].Board, size: this.state.all[index].Size})
            }.bind(this), 3000);
        }
    }
    componentDidMount(){
        this.socketHandler()        
    }
    render(){
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