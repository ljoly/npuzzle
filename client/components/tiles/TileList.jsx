import React, {Component} from 'react';
import PropTypes from 'prop-types'
import Tile from './Tile.jsx'

class TileList extends Component{
    render(){
        return (
            <div className='tile_list'>
                {this.props.tiles.map(tile => {
                        return(<Tile
                        tile={tile}
                        size={this.props.size}
                        key={tile.toString()}
                    />)
                })}
            </div>
        )
    }
}

export default TileList