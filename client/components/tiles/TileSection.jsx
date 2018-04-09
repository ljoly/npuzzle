import React, {Component} from 'react';
import PropTypes from 'prop-types'
import TileList from './TileList.jsx'

class TileSection extends Component{
    render(){
        return (
            <div className='tile_section'>
                {this.props.puzzle.map(tiles => {
                    return (<TileList
                        tiles={tiles}
                        size={this.props.size}
                        {...this.state}
                    />)
                })}
            </div>
        )
    }
}

export default TileSection