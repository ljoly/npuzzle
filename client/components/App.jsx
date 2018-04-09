import React, {Component} from 'react';
import TileSection from './tiles/TileSection.jsx'
import Socket from '../socket.js'

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
    // componentDidMount(){
    //     let socket = this.socket = new Socket();
    //     socket.on('connect', this.onConnect.bind(this));
    //     socket.on('disconnect', this.onDisconnect.bind(this));
    //     socket.on('update', this.updatePuzzle.bind(this));
    // }
    // updatePuzzle(puzzle){
    //     this.setState({puzzle});
    // }
    // onConnect(){
    //     this.setState({connected: true});
    // }
    // onDisconnect(){
    //     this.setState({connected: false});
    // }
    render(){
        return (
            <div>
                <TileSection
                    {...this.state}
                />
            </div>
        )
    }
}

export default App