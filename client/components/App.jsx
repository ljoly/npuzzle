import React, {Component} from 'react';
import TileSection from './tiles/TileSection.jsx'
import socketIOClient from 'socket.io-client'

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            puzzle: [[1, 2, 3, 4, 5],
                    [16, 17, 18, 19, 6],
                    [15, 24, 0, 20, 7],
                    [14, 23, 22, 21, 8],
                    [13, 12, 11, 10, 9]],
            size: 5,
            endpoint: "http://{your machine's ip}:4001" // this is where we are connecting to with sockets
        }
    }
    send() {
        const socket = socketIOClient(this.state.endpoint)
        
        // this emits an event to the socket (your server) with an argument of 'red'
        // you can make the argument any color you would like, or any kind of data you want to send.
        
        socket.emit('hello from front', 'red') 
        // socket.emit('change color', 'red', 'yellow') | you can have multiple arguments
    }
    render(){
        const socket = socketIOClient(this.state.endpoint)
        socket.on('change color', (color) => {
            // setting the color of our button
            document.body.style.backgroundColor = color
          })
        return (
            // <div>
            //     <TileSection
            //         {...this.state}
            //     />
            // </div>
            <div style={{ textAlign: "center" }}>
            <button onClick={() => this.send()}>Change Color</button>
          </div>
        )
    }
}

export default App