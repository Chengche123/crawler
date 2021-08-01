package download

import (
	"fmt"
	"sync"
)

type Repository interface {
	FindStaticUrls() (<-chan []string, error)
}

type Downloader struct {
	Repository Repository
}

func (d *Downloader) StartDownload() error {
	inchan, err := d.Repository.FindStaticUrls()
	if err != nil {
		return fmt.Errorf("failed to find models in dao: %v", err)
	}

	const concur = 64
	var wg sync.WaitGroup

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for urls := range inchan {
				err := DownloadStatic(urls)
				if err != nil {
					// log啥?
					continue
				}

			}

		}()
	}

	wg.Wait()

	return nil
}
