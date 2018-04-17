import React, {Component} from 'react';
import PropTypes from 'prop-types'

class Tile extends Component{
    getColor(number){
        var r = (number & 0xFF).toString(16);
        var g = ((number >> 8) & 0xFF).toString(16);
        var b = ((number >> 16) & 0xFF).toString(16);
     
        r = ('0' + r).slice(-2);
        g = ('0' + g).slice(-2);
        b = ('0' + b).slice(-2);
     
        return "#" + r + g + b;
    }
    render(){
        var color = this.getColor((50 / this.props.size) * this.props.tile)
        if (this.props.tile != 0) {
            return (
                <span className="tile" style={{backgroundColor: color}}>
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