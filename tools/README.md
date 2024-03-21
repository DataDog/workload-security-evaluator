# CPU Spiker

`cpuspiker` is a binary that can be used to test [malware detection](https://www.datadoghq.com/blog/cloud-security-malware-detection/) and [crypto mining protection](https://www.datadoghq.com/blog/datadog-csm-threats-cryptomining/) from Datadog [Cloud Security Management](https://www.datadoghq.com/product/cloud-security-management/). 

## Usage

```
wget https://raw.githubusercontent.com/DataDog/workload-security-evaluator/main/tools/cpuspiker-amd64 -O cpuspiker
chmod +x cpuspiker
./cpuspiker -cpu-priority 1
```