import React from 'react';
import ReactDOM from 'react-dom';
import { Sparklines, SparklinesBars, SparklinesLine, SparklinesNormalBand, SparklinesReferenceLine, SparklinesSpots } from 'react-sparklines';

function boxMullerRandom () {
    let phase = false,
        x1, x2, w, z;

    return (function() {

        if (phase = !phase) {
            do {
                x1 = 2.0 * Math.random() - 1.0;
                x2 = 2.0 * Math.random() - 1.0;
                w = x1 * x1 + x2 * x2;
            } while (w >= 1.0);

            w = Math.sqrt((-2.0 * Math.log(w)) / w);
            return x1 * w;
        } else {
            return x2 * w;
        }
    })();
}

function randomData(n = 30) {
    return Array.apply(0, Array(n)).map(boxMullerRandom);
}

const sampleData = randomData(30);
const sampleData100 = randomData(100);

class FacebookChartForOpenData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>Open</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class FacebookChartForHighData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>High</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class FacebookChartForLowData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>Low</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class FacebookChartForCloseData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>Close</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class FacebookChartForVolumeData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>Volume</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class FacebookChartForAdjData extends React.Component {
    state = {
        open : [32.34453, 27.120001, 43.347, 23.56734, 21.908877, 33.445692],
    }
    render() {
        return (
            <div>
                <h4>Adj</h4>
                <Sparklines data={this.state.open} width={200} height={150}>
                    <SparklinesLine color="#1c8cdc"/>
                </Sparklines>
            </div>
        );
    }
}

class Customizable1 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#1c8cdc" />
            </Sparklines>
        );
    }
}

class Customizable2 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#fa7e17" />
            </Sparklines>
        );
    }
}

class Customizable3 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#ea485c" />
            </Sparklines>
        );
    }
}

class Customizable4 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#56b45d" />
            </Sparklines>
        );
    }
}

class Customizable5 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#8e44af" />
            </Sparklines>
        );
    }
}

class Customizable6 extends React.Component {
    render() {
        return (
            <Sparklines data={sampleData} width={300} height={150}>
                <SparklinesLine color="#253e56" />
                <SparklinesSpots />
            </Sparklines>
        );
    }
}

const demos = {
  'facebookChartForOpenData': FacebookChartForOpenData,
  'facebookChartForHighData': FacebookChartForHighData,
  'facebookChartForLowData': FacebookChartForLowData,
  'facebookChartForCloseData': FacebookChartForCloseData,
  'facebookChartForVolumeData': FacebookChartForVolumeData,
  'facebookChartForAdjData': FacebookChartForAdjData,
  'customizable1': Customizable1,
  'customizable2': Customizable2,
  'customizable3': Customizable3,
  'customizable4': Customizable4,
  'customizable5': Customizable5,
  'customizable6': Customizable6
};

for (let d in demos) {
    ReactDOM.render(React.createElement(demos[d]), document.getElementById(d));
}