import React, { Component } from 'react';
import PropTypes from 'prop-types'
import TileList from './TileList.jsx'

class TileSection extends Component {
    getTabs(array, chunk) {
        var ret = [];
        var i, j, temparray;
        for (i = 0, j = array.length; i < j; i += chunk) {
            temparray = array.slice(i, i + chunk);
            ret.push(temparray)
        }
        return ret
    }
    render() {
        var tabs = this.getTabs(this.props.puzzle, this.props.size)
        return (
            <div className='tile_section'>
                {tabs.map(tiles => {
                    return (<TileList
                        tiles={tiles}
                        size={this.props.size}
                        {...this.state}
                        key={tiles.toString()}
                    />)
                })}
            </div>
        )
    }
}

export default TileSection