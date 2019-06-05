# monitor

report information to cloud monitor(currently only support jdcloud.com)

## Install

```sh
go get github.com/lidaobing/monitor
```

## prepare config

prepare a file like following, put it under `$HOME/.lidaobing-monitor.toml`.

```toml
AK = "YOUR JDCLOUD AK" # Your jdcloud ak
SK = "YOUR JDCLOUD SK" # Your jdcloud sk
Namespace = "computers" # namespace of the monitor results

[[metrics]]
Name = "cpu.load1" # name of the metric
Value = "uptime | awk '{print $(NF-2)}' | tr -d ','" # the shell command used to fetch the result

[[metrics]]
Name = "cpu.load5"
Value = "uptime | awk '{print $(NF-1)}' | tr -d ','"

[[metrics]]
Name = "cpu.load15"
Value = "uptime | awk '{print $(NF-1)}' | tr -d ','"

[[metrics]]
Name = "ping.baidu.ms"
Value = "ping -c 1 www.baidu.com  | head -2 | tail -1 | sed -e 's/^.*time=\\(.*\\) ms/\\1/g'"

[[metrics]]
Name = "rx.kB"
Value = "sar -n DEV | grep wlp3s0 | tail -2 | head -1 | awk '{print $6}'"

[[metrics]]
Name = "tx.kB"
Value = "sar -n DEV | grep wlp3s0 | tail -2 | head -1 | awk '{print $7}'"
```

### install to crontab

```crontab
* * * * * $HOME/go/bin/monitor > /dev/null
```
