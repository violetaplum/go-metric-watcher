package repository

import (
	"context"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"sync"
)

type AlertRepository struct {
	rules map[string]*model.AlertRule
	mutex sync.RWMutex
}

func NewAlertRepository() *AlertRepository {
	return &AlertRepository{
		rules: make(map[string]*model.AlertRule),
	}
}

func (r *AlertRepository) Save(ctx context.Context, rule *model.AlertRule) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.rules[rule.ID] = rule
	return nil
}

func (r *AlertRepository) FindByID(ctx context.Context, id string) (*model.AlertRule, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if rule, exists := r.rules[id]; exists {
		return rule, nil
	}

	return nil, nil
}

func (r *AlertRepository) FindAll(ctx context.Context) ([]*model.AlertRule, error) {
	r.mutex.RLock() // RLock(): 읽기잠금    Lock(): 쓰기잠금
	defer r.mutex.RUnlock()

	rules := make([]*model.AlertRule, 0, len(r.rules))
	for _, rule := range r.rules {
		rules = append(rules, rule)
	}
	return rules, nil
}
