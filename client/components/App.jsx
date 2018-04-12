import React, {Component} from 'react';
import TileSection from './tiles/TileSection.jsx'
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
            size: 5,
            endpoint: "http://localhost:3000" // this is where we are connecting to with sockets
        }
    }
    render(){
        const socket = io(this.state.endpoint)
        return (
            // <div>
            //     <TileSection
            //         {...this.state}
            //     />
            // </div>
            <div style={{ textAlign: "center" }}>
                <button onClick={() => console.log("button")}>Change Color</button>
          </div>
        )
    }
}

export default App