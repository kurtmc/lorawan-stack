// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import PropTypes from 'prop-types'
import { getApplicationId, getDeviceId } from '../../../lib/selectors/id'
import MapWidget from '../../../components/map/widget/'

export default class DeviceMap extends React.Component {
  render() {
    const { device } = this.props

    const devIds = device && device.ids
    const devId = getDeviceId(devIds)
    const appId = getApplicationId(devIds)

    const markers =
      device.locations && device.locations.user
        ? [
            {
              position: {
                latitude: device.locations.user.latitude || 0,
                longitude: device.locations.user.longitude || 0,
              },
            },
          ]
        : false

    return (
      <MapWidget
        id="device-map-widget"
        markers={markers}
        link={`/applications/${appId}/devices/${devId}/location`}
      />
    )
  }
}

DeviceMap.propTypes = {
  device: PropTypes.object.isRequired,
}
