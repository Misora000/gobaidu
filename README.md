# Baidu Search API for golang

A simple golang implementation to get the title, snippet, url and image_url from the Baidu search.

## Usage

```
import "github.com/Misora000/gobaidu/search"

search.Search( context, keyword, max_page )
```

## Example

```go
search.Search(context.Background(), "鋼鐵人", 4)
```

Result
```json
{
  "items": [
    {
      "title": "《钢铁人》第1集_高清在线播放_韩剧_热播韩剧网_韩剧TV网",
      "snippet": "韩剧tv网(http://www.hanjutv.com)为您提供《钢铁人》第1集在线播放资源,更多好看高清韩国电视剧,尽在韩剧tv网",
      "url": "http://www.baidu.com/link?url=w3ZNmy6-1lCAdfkW5yh9Z168zXfAcCsRvXMq2ZUQ2JaMJ3yeQ3yGPcrWHE5lbMgo7NsB4EFvfNP1Mea99WTDiq",
      "image_url": "https://dss2.baidu.com/6ONYsjip0QIZ8tyhnq/it/u=2411828420,4195884097\u0026fm=85\u0026app=92\u0026f=JPEG?w=121\u0026h=75\u0026s=2B335D850E723698162D11EA03005013"
    },
    {
      "title": "鋼鐵人 | 玩具人Toy People News",
      "snippet": "最新最快的玩具新聞與詳盡報導,全部都在玩具人 TOY PEOPLE NEWS!... (Spider-Man: Homecoming)中登場的「鋼鐵人 馬克47 」(IRON MAN MARK XLVII)1/6比例的合金...",
      "url": "http://www.baidu.com/link?url=4-cdBo1os5PBIyfx8zQKlvCDesyCbPjVg9GsdV1fTcyz4zmqW6KafgE-_g2UEbJmNTNJYZvDOtKo1bRoeiykyu3TPNWRA2Li_AkgNlMcYVS",
      "image_url": ""
    },
    ...
    {
      "title": "钢铁人-八三夭, 钢铁人MP3下载,歌词下载 - 虾米音乐",
      "snippet": "《钢铁人》演唱者八三夭,所属专辑《大逃杀》;免费在线试听钢铁人,MP3下载、钢铁人歌词下载;更多八三夭相关歌曲推荐,尽在虾米音乐!",
      "url": "http://www.baidu.com/link?url=m_PGreoSc9wofE_pmBX-0hNXD_X0pfRT10SL6cjVsT5eHcOynjcm7u67olC93VZX",
      "image_url": "https://dss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=199873467,3980739620\u0026fm=85\u0026app=92\u0026f=JPEG?w=121\u0026h=75\u0026s=6DD2EA0BD8510ED462F0D50F010070C1"
    },
    {
      "title": "《鋼鐵人3》電影預告、劇照及海報搶先看! | 電影、好萊塢、鋼鐵人...",
      "snippet": "看完帥哥眾星雲集的《復仇者聯盟》 之後,是不是很想念最有魅力的花花公子 ─ 鋼鐵人呢?《鋼鐵人3》(Iron Man 3)即 電影、好萊塢、鋼鐵人3、Iron Man 3...",
      "url": "http://www.baidu.com/link?url=29m2u7BsnnQ6O0fa19tttVVvNXC-XARu4JKw9g0OPvnfJUqcdLFkAZGISWa0-QBHOL5lf66YiDVr-C0brhxJeq",
      "image_url": ""
    }
  ]
}
```

