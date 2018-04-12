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
                <span className="tile" style={{backgroundColor: bgColor}}>
                    {this.props.tile}  
                </span>
            )
        } else {
            return (
                <span className="tile" style={{backgroundColor: 'white'}}>&nbsp;</span>
            )
        }
    }
}

Tile.propTypes = {
    tile: PropTypes.number.isRequired,
    size: PropTypes.number.isRequired
}

export default Tile