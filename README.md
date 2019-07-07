[![Board Status](https://mchirico.visualstudio.com/c7e39c31-ad05-41c9-8b5b-22ac9a522c14/bde6ae0a-fc94-4b19-a70c-2389949f30e8/_apis/work/boardbadge/54cb834f-ff69-4ab9-b365-6625aa2a9081?columnOptions=1)](https://mchirico.visualstudio.com/c7e39c31-ad05-41c9-8b5b-22ac9a522c14/_boards/board/t/bde6ae0a-fc94-4b19-a70c-2389949f30e8/Microsoft.RequirementCategory/)

[![Build Status](https://mchirico.visualstudio.com/ipblock/_apis/build/status/mchirico.ipblock?branchName=master)](https://mchirico.visualstudio.com/ipblock/_build/latest?definitionId=9&branchName=master)



[![codecov](https://codecov.io/gh/mchirico/ipblock/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/ipblock)
# ipblock

## Notes
```

./ipblock rule ./test-fixtures/mail.log

--- SCRIPT ---

#!/bin/bash
mkdir rules
cd rules
curl http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz -o all-zones.tar.gz
tar -xzf all-zones.tar.gz


for FILE in  de mx jp br ar cn es in co ve fr ua sy ie pa dk kr pl gt it ru
do
  echo -e '#!/bin/bash\n' > ${FILE}.sh
  awk '{printf("iptables -A INPUT -s %!!(MISSING)s(MISSING) -j DROP\n",$1)}' "${FILE}.zone"  >> "${FILE}.sh"
  echo -e 'iptables-save | awk '"'"'!seen[$0]++'"'"'|iptables-restore\n' >> "${FILE}.sh"
  chmod 700 "${FILE}.sh"
done





```
[Ref: www.ipdeny.com/ipblocks ](http://www.ipdeny.com/ipblocks)



## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


