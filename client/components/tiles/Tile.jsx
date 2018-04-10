import React, {Component} from 'react';
import PropTypes from 'prop-types'

class Tile extends Component{
    render(){
        var r = (this.props.tile * this.props.size & 0xFF).toString(16);
        var g = ((this.props.tile * this.props.size >> 8) & 0xFF).toString(16);
        var b = ((this.props.tile * this.props.size >> 16) & 0xFF).toString(16);
     
        r = ('0' + r).slice(-2);
        g = ('0' + g).slice(-2);
        b = ('0' + b).slice(-2);
     
        const bgColor = "#" + r + g + b;
        if (this.props.tile != 0) {
            return (
                <tile style={{backgroundColor: bgColor}}>
                    {this.props.tile}  
                </tile>
            )
        } else {
            return (
                <tile style={{backgroundColor: 'white'}}>&nbsp;</tile>
            )
        }
    }
}

Tile.propTypes = {
    tile: PropTypes.object.isRequired,
    size: PropTypes.object.isRequired
}

export default Tile