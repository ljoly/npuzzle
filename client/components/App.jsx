import React, {Component} from 'react';
import TileSection from './tiles/TileSection.jsx'
import ButtonSection from './buttons/ButtonSection.jsx'
import io from 'socket.io-client'

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            puzzle: [[1, 2, 3, 4, 5],
                    [16, 17, 18, 19, 6],
                    [15, 24, 0, 20, 7],
                    [14, 23, 22, 21, 8],
                    [13, 12, 11, 10, 9]],
            size: 5
        }
    }
    prevState(){
        console.log('prev')
    }
    nextState(){
        console.log('next')
    }
    go(){
        console.log('go')
    }
    render(){
        const socket = io("http://localhost:3000")
        if (socket.connected) {
            console.log("CONNECTED TO SERVER")
        } else {
            console.log("CANNOT CONNECT")
        }
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
        //     <div style={{ textAlign: "center" }}>
        //         <button onClick={() => console.log("button")}>Change Color</button>
        //   </div>
        )
    }
}

export default App